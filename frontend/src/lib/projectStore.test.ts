import { expect, test, vi } from 'vitest';
import { frontend } from './wailsjs/go/models';
import { ProjectStore } from './projectStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/Projects';
import ProjectDto = frontend.ProjectDto;

test('create project store', async () => {
	vi.spyOn(models, 'Create').mockImplementation((projectDto: ProjectDto): Promise<ProjectDto> => {
		return Promise.resolve(projectDto);
	});

	vi.spyOn(models, 'Update').mockImplementation((projectDto: ProjectDto): Promise<ProjectDto> => {
		return Promise.resolve(projectDto);
	});

	vi.spyOn(models, 'Delete').mockImplementation((projectDto: ProjectDto): Promise<void> => {
		return Promise.resolve();
	});

	const projectDtos = [];
	const projectDtoOne = new ProjectDto('{"id":1,"updatedAt":"2024-08-28T20:49:26.36341633+02:00","name":"Project 0"}');
	const projectDtoTwo = new ProjectDto('{"id":2,"updatedAt":"2024-08-28T20:49:26.36341633+02:00","name":"Project 1"}');
	projectDtos.push(projectDtoOne, projectDtoTwo);
	const projectStore = new ProjectStore(projectDtos);
	expect(projectStore.projects[0]).toBe(projectDtoOne);
	expect(projectStore.projects[1]).toBe(projectDtoTwo);

	projectStore.create('yolo');
	await vi.waitFor(() => {
		expect(projectStore.projects[0].name).toBe('yolo');
	});

	projectStore.projects[0].name = 'zwolo';
	projectStore.update(projectStore.projects[0]);
	await vi.waitFor(() => {
		expect(projectStore.projects[0].name).toBe('zwolo');
	});

	projectStore.delete(projectStore.projects[0]);
	await vi.waitFor(() => {
		expect(projectStore.projects.length).toBe(2);
	});

	expect(projectStore.getById(projectDtoTwo.id).name).toBe('Project 1');
});
