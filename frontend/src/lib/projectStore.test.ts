import { expect, test } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { ProjectStore } from './projectStore.svelte';
import ProjectDto = frontend.ProjectDto;

test('create project store', () => {
	const projectDtos = [];
	const projectDtoOne = new ProjectDto('{"id":1,"updatedAt":"2024-08-28T20:49:26.36341633+02:00","name":"Project 0"}');
	const projectDtoTwo = new ProjectDto('{"id":2,"updatedAt":"2024-08-28T20:49:26.36341633+02:00","name":"Project 1"}');
	projectDtos.push(projectDtoOne, projectDtoTwo);
	const projectStore = new ProjectStore(projectDtos);
	expect(projectStore.projects[0]).toBe(projectDtoOne);
	expect(projectStore.projects[1]).toBe(projectDtoTwo);
});
