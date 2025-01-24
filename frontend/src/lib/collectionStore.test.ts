import { expect, test, vi } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { CollectionStore } from './collectionStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/Collections';
import CollectionDto = frontend.CollectionDto;

test('create collection store', async () => {
	vi.spyOn(models, 'Create').mockImplementation((collectionDto: CollectionDto): Promise<CollectionDto> => {
		collectionDto.id = 3;

		return Promise.resolve(collectionDto);
	});

	vi.spyOn(models, 'Update').mockImplementation((collectionDto: CollectionDto): Promise<CollectionDto> => {
		return Promise.resolve(collectionDto);
	});

	vi.spyOn(models, 'Delete').mockImplementation((collectionDto: CollectionDto): Promise<void> => {
		return Promise.resolve();
	});

	const collectionDtos = [];
	const collectionDtoOne = new CollectionDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.366552118+02:00","name":"Collection 0","projectId":1}'
	);
	const collectionDtoTwo = new CollectionDto(
		'{"id":2,"updatedAt":"2024-08-28T20:49:26.366552118+02:00","name":"Collection 1","projectId":2}'
	);
	collectionDtos.push(collectionDtoOne, collectionDtoTwo);
	const collectionStore = new CollectionStore(collectionDtos);
	expect(collectionStore.getByProjectId(1)[0]).toBe(collectionDtoOne);
	expect(collectionStore.getByProjectId(2)[0]).toBe(collectionDtoTwo);

	collectionStore.create(1, 'yolo');
	await vi.waitFor(() => {
		expect(collectionStore.getById(3).name).toBe('yolo');
	});

	collectionStore.getById(3).name = 'zwolo';
	collectionStore.update(collectionStore.getById(3));
	await vi.waitFor(() => {
		expect(collectionStore.getById(3).name).toBe('zwolo');
	});

	collectionStore.delete(collectionStore.getById(3));
	await vi.waitFor(() => {
		expect(collectionStore.getByProjectId(1).length).toBe(1);
	});
});
