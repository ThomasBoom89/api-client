import { expect, test } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { CollectionStore } from './collectionStore.svelte';
import CollectionDto = frontend.CollectionDto;

test('create collection store', () => {
	const collectionDtos = [];
	const collectionDtoOne = new CollectionDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.366552118+02:00","name":"Collection 0","projectId":1}',
	);
	const collectionDtoTwo = new CollectionDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.366552118+02:00","name":"Collection 0","projectId":2}',
	);
	collectionDtos.push(collectionDtoOne, collectionDtoTwo);
	const collection = new CollectionStore(collectionDtos);
	expect(collection.getByProjectId(1)[0]).toBe(collectionDtoOne);
	expect(collection.getByProjectId(2)[0]).toBe(collectionDtoTwo);
});
