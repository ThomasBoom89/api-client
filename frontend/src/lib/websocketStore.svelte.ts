import { getContext, setContext } from 'svelte';
import { EventsOn } from '$lib/wailsjs/runtime';
import { frontend } from '$lib/wailsjs/go/models.ts';
import { Connect, Disconnect, Send } from '$lib/wailsjs/go/frontend/Websocket';
import { SvelteMap } from 'svelte/reactivity';
import { WebsocketMessageType } from '$lib/enums/WebsocketMessageType.ts';

class WebsocketMessage {
	public value: string = $state('');
	public type: WebsocketMessageType = $state(WebsocketMessageType.Incoming);
}

class InternalState {
	public messages: WebsocketMessage[] = $state([]);
	public error: string = $state('');
	public connected: boolean = $state(false);
}

export class WebsocketStore {
	public connections: SvelteMap<number, InternalState> = $state(new SvelteMap());
	public event;

	public constructor() {
		this.event = EventsOn('websocket', (websocketStateDto: frontend.WebsocketStateDto) => {
			let internalState = this.connections.get(websocketStateDto.id);
			if (internalState === undefined) {
				return;
			}
			if (websocketStateDto.incomingMessage !== '') {
				let incomingMessage = new WebsocketMessage();
				incomingMessage.type = WebsocketMessageType.Incoming;
				incomingMessage.value = websocketStateDto.incomingMessage;
				internalState.messages.reverse();
				internalState.messages.push(incomingMessage);
				internalState.messages.reverse();
			}
			internalState.error = websocketStateDto.error;
			internalState.connected = websocketStateDto.connected;
		});
	}

	public connect(websocketRequestDto: frontend.WebsocketRequestDto): void {
		Connect(websocketRequestDto).then((response: frontend.WebsocketStateDto) => {
			if (this.connections.get(websocketRequestDto.id) === undefined) {
				this.connections.set(websocketRequestDto.id, new InternalState());
			}
			// @ts-ignore
			this.connections.get(websocketRequestDto.id).connected = response.connected;
			// @ts-ignore
			this.connections.get(websocketRequestDto.id).error = response.error;
		});
	}

	public sendMessage(websocketRequestDto: frontend.WebsocketRequestDto, message: string): void {
		Send(websocketRequestDto, message).then((websocketStateDto: frontend.WebsocketStateDto) => {
			let internalState = this.connections.get(websocketRequestDto.id);
			if (internalState === undefined || !websocketStateDto.connected) {
				return;
			}
			internalState.error = websocketStateDto.error;
			internalState.connected = websocketStateDto.connected;
			let outgoingMessage = new WebsocketMessage();
			outgoingMessage.type = WebsocketMessageType.Outgoing;
			outgoingMessage.value = message;
			internalState.messages.reverse();
			internalState.messages.push(outgoingMessage);
			internalState.messages.reverse();
		});
	}

	public getStateById(requestID: number): InternalState {
		const state = this.connections.get(requestID);
		if (state === undefined) {
			return new InternalState();
		}

		return state;
	}

	public disconnect(websocketRequestDto: frontend.WebsocketRequestDto): void {
		Disconnect(websocketRequestDto).then((websocketStateDto: frontend.WebsocketStateDto) => {
			let internalState = this.connections.get(websocketRequestDto.id);
			if (internalState === undefined) {
				return;
			}
			internalState.error = websocketStateDto.error;
			internalState.connected = websocketStateDto.connected;
		});
	}
}

const websocketStoreContextKey = 'websocketStore';

export function initializeWebsocketStore(): void {
	setContext<WebsocketStore>(websocketStoreContextKey, new WebsocketStore());
}

export function getWebsocketStore(): WebsocketStore {
	return getContext<WebsocketStore>(websocketStoreContextKey);
}
