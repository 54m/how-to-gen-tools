// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 2.3.1

export type Meta = {
	createdAt: string;
	createdBy: string;
	deletedAt: string | null;
	deletedBy: string;
	updatedAt: string;
	updatedBy: string;
	version: number;
}
export type User = {
	age: number;
	createdAt: string;
	createdBy: string;
	deletedAt: string | null;
	deletedBy: string;
	id: string;
	name: string;
	updatedAt: string;
	updatedBy: string;
	version: number;
}
export type FailedReason = {
	message: string;
}
export type DeleteUserRequest = {
	userID: string;
}
export type DeleteUserResponse = {
	messages?: FailedReason[];
	status: number;
}
export type GetRequest = {
	age?: number;
	limit?: number;
	name?: string;
	pagingKey?: string;
	sortOrder?: string;
}
export type GetResponse = {
	messages?: FailedReason[];
	nextPagingKey?: string;
	payload: (User | null)[] | null;
	status: number;
}
export type GetUserRequest = {
	userID: string;
}
export type GetUserResponse = {
	messages?: FailedReason[];
	payload?: User;
	status: number;
}
export type PostRequest = {
	age: number;
	name: string;
}
export type PostResponse = {
	messages?: FailedReason[];
	payload?: User;
	status: number;
}
export type PutUserRequest = {
	age: number;
	name: string;
	userID: string;
	version: number;
}
export type PutUserResponse = {
	messages?: FailedReason[];
	payload?: User;
	status: number;
}
export type Pagination = {
	limit?: number;
	pagingKey?: string;
	sortOrder?: string;
}
