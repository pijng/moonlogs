import { BaseResponse, del, get, post, put } from "./base";

export type NotificationMethod = "POST" | "GET";

export type NotificationHeader = {
  key: string;
  value: string;
};

export type NotificationProfile = {
  id: number;
  name: string;
  rule_ids: number[];
  description: string;
  enabled: boolean;
  silence_for: string;
  url: string;
  method: NotificationMethod;
  headers: NotificationHeader[];
  payload: string;
};

export type NotificationProfileToCreate = {
  name: string;
  description: string;
  rule_ids: number[];
  enabled: boolean;
  silence_for: string;
  url: string;
  method: NotificationMethod;
  headers: NotificationHeader[];
  payload: string;
};

export type NotificationProfileToUpdate = {
  id: number;
  name: string;
  rule_ids: number[];
  description: string;
  enabled: boolean;
  silence_for: string;
  url: string;
  method: NotificationMethod;
  headers: NotificationHeader[];
  payload: string;
};

export interface NotificationProfileResponse extends BaseResponse {
  data: NotificationProfile;
}

export interface NotificationProfileListResponse extends BaseResponse {
  data: NotificationProfile[];
}

export const getNotificationProfiles = (): Promise<NotificationProfileListResponse> => {
  return get({ url: "/api/notification_profiles" });
};

export const getNotificationProfile = (id: number): Promise<NotificationProfileResponse> => {
  return get({ url: `/api/notification_profiles/${id}` });
};

export const createNotificationProfile = (profile: NotificationProfileToCreate): Promise<NotificationProfileResponse> => {
  return post({ url: "/api/notification_profiles", body: JSON.stringify(profile) });
};

export const editNotificationProfile = (profile: NotificationProfileToUpdate): Promise<NotificationProfileResponse> => {
  return put({ url: `/api/notification_profiles/${profile.id}`, body: JSON.stringify(profile) });
};

export const deleteNotificationProfile = (id: number): Promise<NotificationProfileResponse> => {
  return del({ url: `/api/notification_profiles/${id}` });
};
