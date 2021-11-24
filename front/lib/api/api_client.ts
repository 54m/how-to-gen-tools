// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 2.3.1

import {
	GetHealthCheckRequest as ApiGetHealthCheckRequest,
	GetHealthCheckResponse as ApiGetHealthCheckResponse,
} from './classes/api/types';

import {
	DeleteUserRequest as ApiUsersDeleteUserRequest,
	DeleteUserResponse as ApiUsersDeleteUserResponse,
	GetRequest as ApiUsersGetRequest,
	GetResponse as ApiUsersGetResponse,
	GetUserRequest as ApiUsersGetUserRequest,
	GetUserResponse as ApiUsersGetUserResponse,
	PostRequest as ApiUsersPostRequest,
	PostResponse as ApiUsersPostResponse,
	PutUserRequest as ApiUsersPutUserRequest,
	PutUserResponse as ApiUsersPutUserResponse,
} from './classes/api/users/types';

export interface MiddlewareContext {
	httpMethod: string;
	endpoint: string;
	request: unknown;
	response?: unknown;
	baseURL: string;
	headers: {[key: string]: string};
	options: {[key: string]: any};
}

export enum MiddlewareResult {
	CONTINUE = 1,
	MIDDLEWARE_STOP = 2,
	STOP = 4
}

export type ApiClientMiddlewareFunc = (context: MiddlewareContext) => Promise<MiddlewareResult|boolean>;

export interface middlewareSet {
	beforeMiddleware?: ApiClientMiddlewareFunc[];
	afterMiddleware?: ApiClientMiddlewareFunc[];
}

class ApiClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
	public users: ApiUsersClient;
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};
		this.users = new ApiUsersClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
	}

	getRequestObject(obj: any, routingPath: string[]): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (routingPath.indexOf(key) === -1) {
				res[key] = obj[key];
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}

	async getHealthCheck(
		param: ApiGetHealthCheckRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiGetHealthCheckResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/api/health_check`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/health_check?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiGetHealthCheckResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class ApiUsersClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
	}

	getRequestObject(obj: any, routingPath: string[]): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (routingPath.indexOf(key) === -1) {
				res[key] = obj[key];
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}

	async deleteUser(
		param: ApiUsersDeleteUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiUsersDeleteUserResponse> {
	    const excludeParams: string[] = ['userID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'DELETE',
			endpoint: `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}`;
		const resp = await fetch(
			url,
			{
				method: "DELETE",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiUsersDeleteUserResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}

	async get(
		param: ApiUsersGetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiUsersGetResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/api/users/`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/users/?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiUsersGetResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}

	async getUser(
		param: ApiUsersGetUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiUsersGetUserResponse> {
	    const excludeParams: string[] = ['userID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'GET',
			endpoint: `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString();
		const resp = await fetch(
			url,
			{
				method: "GET",
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiUsersGetUserResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}

	async post(
		param: ApiUsersPostRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiUsersPostResponse> {
	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/api/users/`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/users/`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiUsersPostResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}

	async putUser(
		param: ApiUsersPutUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ApiUsersPutUserResponse> {
	    const excludeParams: string[] = ['userID', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'PUT',
			endpoint: `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}`,
			request: param,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/api/users/${encodeURI(param.userID.toString())}`;
		const resp = await fetch(
			url,
			{
				method: "PUT",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ApiUsersPutUserResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];

	public api: ApiClient;

	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions: {[key: string]: any} = {},
		middleware?: middlewareSet
	) {
		const headers: {[key: string]: string} =  {
			'Content-Type': 'application/json',
			...commonHeaders,
		};

		if (token !== undefined) {
			headers['Authorization'] = 'Bearer ' + token;
		}

		this.baseURL =  (baseURL === undefined) ? "" : baseURL;
		this.options = commonOptions;
		this.headers = headers;

		if (middleware) {
			this.beforeMiddleware = middleware.beforeMiddleware ?? [];
			this.afterMiddleware = middleware.afterMiddleware ?? [];
		}
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};

		this.api = new ApiClient(
			headers,
			this.options,
			this.baseURL,
			childMiddlewareSet
		);
	}

	getRequestObject(obj: any, routingPath: string[]): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (routingPath.indexOf(key) === -1) {
				res[key] = obj[key];
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}
}

export class ApiError extends Error {
	private _statusCode: number;
	private _statusText: string;
	private _response: string;

	constructor(response: Response, responseText: string) {
		super();
		this._statusCode = response.status;
		this._statusText = response.statusText;
		this._response = responseText
	}

	get statusCode(): number {
		return this._statusCode;
	}

	get statusText(): string {
		return this._statusText;
	}

	get response(): string {
		return this._response;
	}
}

export class ApiMiddlewareStop extends Error {}

export interface MockOption {
	wait_ms: number;
	target_file: string;
}
