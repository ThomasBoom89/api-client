import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Create, Delete, Update } from './wailsjs/go/frontend/Projects';
import ProjectDto = frontend.ProjectDto;

export class ProjectStore {
	private _projects: ProjectDto[] = $state([]);

	get projects(): ProjectDto[] {
		return this._projects;
	}

	constructor(projects: ProjectDto[]) {
		this._projects = projects;
	}

	public create(name: string): void {
		let dto = new ProjectDto();
		dto.name = name;
		Create(dto).then((project) => {
			let projectsReversed = this._projects.toReversed();
			projectsReversed.push(project);
			this._projects = projectsReversed.toReversed();
		});
	}

	public delete(project: ProjectDto): void {
		Delete(project).then(() => {
			this._projects = this._projects.filter((_project) => {
				return project.id !== _project.id;
			});
		});
	}

	public update(project: ProjectDto): void {
		Update(project).then((project) => {
			let index = this._projects.findIndex((_project) => _project.id === project.id);
			this._projects[index] = project;
		});
	}
}

const projectStoreContextKey = 'projectStore';

export function initializeProjectStore(projects: ProjectDto[]): void {
	setContext<ProjectStore>(projectStoreContextKey, new ProjectStore(projects));
}

export function getProjectStore(): ProjectStore {
	return getContext<ProjectStore>(projectStoreContextKey);
}
