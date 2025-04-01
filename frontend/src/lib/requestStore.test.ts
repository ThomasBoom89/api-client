import { expect, test, vi } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { RequestStore } from './requestStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/HttpRequests';
import { RequestTypes } from '$lib/enums/RequestTypes';
import HttpRequestDto = frontend.HttpRequestDto;

test('create request store', async () => {
	vi.spyOn(models, 'Create').mockImplementation((httpRequestDto: HttpRequestDto): Promise<HttpRequestDto> => {
		httpRequestDto.id = 3;
		httpRequestDto.type = RequestTypes.HTTP;

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
	const requestStore = new RequestStore(requestDtos);
	expect(requestStore.getByCollectionId(1)[0]).toBe(requestDtoOne);
	expect(requestStore.getByCollectionId(2)[0]).toBe(requestDtoTwo);

	requestStore.create(1, 'yolo', RequestTypes.HTTP);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1)[0].name).toBe('yolo');
	});

	requestStore.getByCollectionId(1)[0].name = 'zwolo';
	requestStore.update(requestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1)[0].name).toBe('zwolo');
	});
	requestStore.delete(requestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1).length).toBe(1);
	});
});
