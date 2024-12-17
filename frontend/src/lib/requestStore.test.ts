import { expect, test, vi } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { RequestStore } from './requestStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/HttpRequests';
import HttpRequestDto = frontend.HttpRequestDto;

test('create request store', async () => {
	vi.spyOn(models, 'Create').mockImplementation((httpRequestDto: HttpRequestDto): Promise<HttpRequestDto> => {
		httpRequestDto.id = 3;

		return Promise.resolve(httpRequestDto);
	});

	vi.spyOn(models, 'Update').mockImplementation((httpRequestDto: HttpRequestDto): Promise<HttpRequestDto> => {
		return Promise.resolve(httpRequestDto);
	});

	vi.spyOn(models, 'Delete').mockImplementation((httpRequestDto: HttpRequestDto): Promise<void> => {
		return Promise.resolve();
	});

	const requestDtos = [];
	const requestDtoOne = new HttpRequestDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest0","type":"http","collectionId":1,"url":"url0"}',
	);
	const requestDtoTwo = new HttpRequestDto(
		'{"id":2,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest1","type":"http","collectionId":2,"url":"url1"}',
	);
	requestDtos.push(requestDtoOne, requestDtoTwo);
	const httpRequestStore = new RequestStore(requestDtos);
	expect(httpRequestStore.getByCollectionId(1)[0]).toBe(requestDtoOne);
	expect(httpRequestStore.getByCollectionId(2)[0]).toBe(requestDtoTwo);

	httpRequestStore.create(1, 'yolo');
	await vi.waitFor(() => {
		expect(httpRequestStore.getByCollectionId(1)[0].name).toBe('yolo');
	});

	httpRequestStore.getByCollectionId(1)[0].name = 'zwolo';
	httpRequestStore.update(httpRequestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(httpRequestStore.getByCollectionId(1)[0].name).toBe('zwolo');
	});

	httpRequestStore.delete(httpRequestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(httpRequestStore.getByCollectionId(1).length).toBe(1);
	});
});
