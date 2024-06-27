export namespace configuration {
	
	export class Configuration {
	
	
	    static createFrom(source: any = {}) {
	        return new Configuration(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

