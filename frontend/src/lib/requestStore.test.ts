import { expect, test } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { RequestStore } from './requestStore.svelte';
import HttpRequestDto = frontend.HttpRequestDto;

test('create request store', () => {
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
});
