import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Create, Update, Delete } from './wailsjs/go/frontend/HttpRequests';
import HttpRequestDto = frontend.HttpRequestDto;

export class RequestStore {
	private _requests: HttpRequestDto[] = $state([]);

	constructor(requests: HttpRequestDto[]) {
		this._requests = requests;
	}

	public getByCollectionId(collectionId: number): HttpRequestDto[] {
		return this._requests.filter((request: HttpRequestDto) => request.collectionId === collectionId);
	}

	public update(request: HttpRequestDto): void {
		Update(request).then((newRequest: HttpRequestDto) => {
			const index = this._requests.findIndex((request) => request.id === newRequest.id);
			this._requests[index] = newRequest;
		});
	}

	public create(collectionId: number, newRequestName: string): void {
		let requestDto = new HttpRequestDto();
		requestDto.collectionId = collectionId;
		requestDto.name = newRequestName;
		Create(requestDto).then((newRequest: HttpRequestDto) => {
			let collections = this._requests.toReversed();
			collections.push(newRequest);
			this._requests = collections.toReversed();
		});
	}

	public delete(requestDto: HttpRequestDto): void {
		Delete(requestDto).then(() => {
			this._requests = this._requests.filter((request) => request.id !== requestDto.id);
		});
	}
}

const requestStoreContextKey = 'requestStore';

export function initializeRequestStore(requests: HttpRequestDto[]): void {
	setContext<RequestStore>(requestStoreContextKey, new RequestStore(requests));
}

export function getRequestStore(): RequestStore {
	return getContext<RequestStore>(requestStoreContextKey);
}
