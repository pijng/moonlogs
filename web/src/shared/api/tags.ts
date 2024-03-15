import { BaseResponse, del, get, post, put } from "./base";

export type Tag = {
  id: number;
  name: string;
  view_order: number;
};

export type TagToCreate = {
  name: string;
  view_order: number | string;
};

export type TagToUpdate = {
  id: number;
  name: string;
  view_order: number | string;
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
  const view_order = tag.view_order === "" ? 0 : tag.view_order;
  const modifiedTag: TagToCreate = { ...tag, view_order: view_order };

  return post({ url: "/api/tags", body: JSON.stringify(modifiedTag) });
};

export const editTag = (tag: TagToUpdate): Promise<TagResponse> => {
  const view_order = tag.view_order === "" ? 0 : tag.view_order;
  const modifiedTag: TagToUpdate = { ...tag, view_order: view_order };

  return put({ url: `/api/tags/${tag.id}`, body: JSON.stringify(modifiedTag) });
};

export const deleteTag = (id: number): Promise<TagResponse> => {
  return del({ url: `/api/tags/${id}` });
};
