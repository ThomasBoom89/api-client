import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Update } from './wailsjs/go/frontend/HttpRequests';
import HttpRequestDto = frontend.HttpRequestDto;

export class RequestStore {
	private readonly _requests: HttpRequestDto[] = $state([]);

	constructor(requests: HttpRequestDto[]) {
		this._requests = requests;
	}

	public getByCollectionId(collectionId: number): HttpRequestDto[] {
		return this._requests.filter((request: HttpRequestDto) => request.collectionId === collectionId);
	}

	public update(request: HttpRequestDto): Promise<HttpRequestDto> {
		return Update(request);
	}
}

const requestStoreContextKey = 'requestStore';

export function initializeRequestStore(requests: HttpRequestDto[]): void {
	setContext<RequestStore>(requestStoreContextKey, new RequestStore(requests));
}

export function getRequestStore(): RequestStore {
	return getContext<RequestStore>(requestStoreContextKey);
}
