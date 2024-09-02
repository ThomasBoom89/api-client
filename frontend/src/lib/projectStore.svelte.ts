import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import ProjectDto = frontend.ProjectDto;

export class ProjectStore {
	private readonly _projects: ProjectDto[] = $state([]);

	get projects(): ProjectDto[] {
		return this._projects;
	}

	constructor(projects: ProjectDto[]) {
		this._projects = projects;
	}
}

const projectStoreContextKey = 'projectStore';

export function initializeProjectStore(projects: ProjectDto[]): void {
	setContext<ProjectStore>(projectStoreContextKey, new ProjectStore(projects));
}

export function getProjectStore(): ProjectStore {
	return getContext<ProjectStore>(projectStoreContextKey);
}
