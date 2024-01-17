import { BaseResponse, del, get, post, put } from "./base";

export type Tag = {
  id: number;
  name: string;
};

export type TagToCreate = {
  name: string;
};

export type TagToUpdate = {
  id: number;
  name: string;
};

export interface TagResponse extends BaseResponse {
  data: Tag;
}

export interface TagListResponse extends BaseResponse {
  data: Tag[];
}

export const getTags = (): Promise<TagListResponse> => {
  return get({ url: "/api/tags" });
};

export const getTag = (id: number): Promise<TagResponse> => {
  return get({ url: `/api/tags/${id}` });
};

export const createTag = (tag: TagToCreate): Promise<TagResponse> => {
  return post({ url: "/api/tags", body: JSON.stringify(tag) });
};

export const editTag = (tag: TagToUpdate): Promise<TagResponse> => {
  return put({ url: `/api/tags/${tag.id}`, body: JSON.stringify(tag) });
};

export const deleteTag = (id: number): Promise<TagResponse> => {
  return del({ url: `/api/tags/${id}` });
};
