export namespace configuration {
	
	export class Configuration {
	    theme: string;
	
	    static createFrom(source: any = {}) {
	        return new Configuration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.theme = source["theme"];
	    }
	}

}

export namespace frontend {
	
	export class CollectionDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    name: string;
	    projectId: number;
	
	    static createFrom(source: any = {}) {
	        return new CollectionDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.name = source["name"];
	        this.projectId = source["projectId"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class HttpRequestBodyDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    httpRequestID: number;
	    type: string;
	    payload: string;
	
	    static createFrom(source: any = {}) {
	        return new HttpRequestBodyDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.httpRequestID = source["httpRequestID"];
	        this.type = source["type"];
	        this.payload = source["payload"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class HttpRequestHeaderDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    httpRequestID: number;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new HttpRequestHeaderDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.httpRequestID = source["httpRequestID"];
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class HttpRequestParameterDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    httpRequestID: number;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new HttpRequestParameterDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.httpRequestID = source["httpRequestID"];
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class HttpRequestDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    name: string;
	    type: string;
	    collectionId: number;
	    url: string;
	    method: string;
	    body: HttpRequestBodyDto;
	    parameter: HttpRequestParameterDto[];
	    header: HttpRequestHeaderDto[];
	
	    static createFrom(source: any = {}) {
	        return new HttpRequestDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.collectionId = source["collectionId"];
	        this.url = source["url"];
	        this.method = source["method"];
	        this.body = this.convertValues(source["body"], HttpRequestBodyDto);
	        this.parameter = this.convertValues(source["parameter"], HttpRequestParameterDto);
	        this.header = this.convertValues(source["header"], HttpRequestHeaderDto);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	
	export class ProjectDto {
	    id: number;
	    // Go type: time
	    updatedAt: any;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ProjectDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	        this.name = source["name"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RequestResponseDTO {
	    error: string;
	    url: string;
	    method: string;
	    responseBody: string;
	    sendHeader: Record<string, string[]>;
	    receivedHeader: Record<string, string[]>;
	    elapsedTime: string;
	    statusCode: number;
	
	    static createFrom(source: any = {}) {
	        return new RequestResponseDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.url = source["url"];
	        this.method = source["method"];
	        this.responseBody = source["responseBody"];
	        this.sendHeader = source["sendHeader"];
	        this.receivedHeader = source["receivedHeader"];
	        this.elapsedTime = source["elapsedTime"];
	        this.statusCode = source["statusCode"];
	    }
	}

}

