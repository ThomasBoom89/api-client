import { expect, test, vi } from 'vitest';
import { frontend } from './wailsjs/go/models';
import * as models from '$lib/wailsjs/go/frontend/Websocket';
import * as events from '$lib/wailsjs/runtime/runtime';
import WebsocketRequestDto = frontend.WebsocketRequestDto;
import WebsocketStateDto = frontend.WebsocketStateDto;
import { WebsocketStore } from '$lib/websocketStore.svelte.ts';

test('create request store', async () => {
	vi.spyOn(events, 'EventsOn').mockImplementation(
		(eventName: string, callback: (...data: any) => void): (() => void) => {
			const websocketStateDto = new WebsocketStateDto(
				'{"id":1,"error":"","connected":true,"incomingMessage":"hallo incoming"}',
			);
			document.addEventListener('websocket', () => {
				callback(websocketStateDto);
			});

			return () => callback(websocketStateDto);
		},
	);

	vi.spyOn(models, 'Connect').mockImplementation(
		(websocketRequestDto: WebsocketRequestDto): Promise<WebsocketStateDto> => {
			let websocketStateDto = new WebsocketStateDto('{"id":1,"error":"","connected":true,"incomingMessage":""}');
			websocketStateDto.id = websocketRequestDto.id;

			return Promise.resolve(websocketStateDto);
		},
	);

	vi.spyOn(models, 'Disconnect').mockImplementation(
		(websocketRequestDto: WebsocketRequestDto): Promise<WebsocketStateDto> => {
			let websocketStateDto = new WebsocketStateDto('{"id":1,"error":"","connected":false,"incomingMessage":""}');
			websocketStateDto.id = websocketRequestDto.id;

			return Promise.resolve(websocketStateDto);
		},
	);

	vi.spyOn(models, 'Send').mockImplementation(
		(websocketRequestDto: WebsocketRequestDto, message: string): Promise<WebsocketStateDto> => {
			let websocketStateDto = new WebsocketStateDto('{"id":1,"error":"","connected":true,"incomingMessage":""}');
			websocketStateDto.id = websocketRequestDto.id;

			return Promise.resolve(websocketStateDto);
		},
	);

	const requestStore = new WebsocketStore();
	const requestDtoOne = new WebsocketRequestDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest0","type":"websocket","collectionId":1,"url":"url0"}',
	);
	requestStore.connect(requestDtoOne);
	await vi.waitFor(() => {
		expect(requestStore.connections.get(requestDtoOne.id)?.connected).is.true;
	});
	requestStore.sendMessage(requestDtoOne, 'hallo');
	await vi.waitFor(() => {
		expect(requestStore.getStateById(requestDtoOne.id)?.messages[0].value).equal('hallo');
	});

	const event = new Event('websocket');
	document.dispatchEvent(event);
	await vi.waitFor(() => {
		expect(requestStore.getStateById(requestDtoOne.id)?.messages[0].value).equal('hallo incoming');
	});

	requestStore.disconnect(requestDtoOne);
	await vi.waitFor(() => {
		expect(requestStore.getStateById(requestDtoOne.id)?.connected).is.false;
	});
});
