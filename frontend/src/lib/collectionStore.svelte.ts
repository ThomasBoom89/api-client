import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import CollectionDto = frontend.CollectionDto;

export class CollectionStore {
	private readonly _collections: CollectionDto[] = $state([]);

	constructor(collections: CollectionDto[]) {
		this._collections = collections;
	}

	public getByProjectId(projectId: number): CollectionDto[] {
		return this._collections.filter((collection) => collection.projectId === projectId);
	}
}

const collectionStoreContextKey = 'collectionStore';

export function initializeCollectionStore(collections: CollectionDto[]): void {
	setContext<CollectionStore>(collectionStoreContextKey, new CollectionStore(collections));
}

export function getCollectionStore(): CollectionStore {
	return getContext<CollectionStore>(collectionStoreContextKey);
}
