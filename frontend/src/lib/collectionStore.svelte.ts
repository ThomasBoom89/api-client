import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import CollectionDto = frontend.CollectionDto;
import { Create, Delete, Update } from './wailsjs/go/frontend/Collections';

export class CollectionStore {
	private _collections: CollectionDto[] = $state([]);

	constructor(collections: CollectionDto[]) {
		this._collections = collections;
	}

	public getByProjectId(projectId: number): CollectionDto[] {
		return this._collections.filter((collection) => collection.projectId === projectId);
	}

	public create(projectId: number, newCollectionName: string): void {
		let collectionDTO = new CollectionDto();
		collectionDTO.projectId = projectId;
		collectionDTO.name = newCollectionName;
		Create(collectionDTO).then((newCollection: CollectionDto) => {
			let collections = this._collections.toReversed();
			collections.push(newCollection);
			this._collections = collections.toReversed();
		});
	}

	public delete(collectionDto: CollectionDto): void {
		Delete(collectionDto).then(() => {
			this._collections = this._collections.filter((collection) => collection.id !== collectionDto.id);
		});
	}

	public update(collectionDto: CollectionDto): void {
		Update(collectionDto).then((collection) => {
			const index = this._collections.findIndex((collection) => collection.id === collectionDto.id);
			this._collections[index] = collection;
		});
	}

	public getById(id: number): CollectionDto {
		return this._collections.find((collection) => collection.id === id) || new CollectionDto();
	}
}

const collectionStoreContextKey = 'collectionStore';

export function initializeCollectionStore(collections: CollectionDto[]): void {
	setContext<CollectionStore>(collectionStoreContextKey, new CollectionStore(collections));
}

export function getCollectionStore(): CollectionStore {
	return getContext<CollectionStore>(collectionStoreContextKey);
}
