import { attach, createEffect } from "effector";
import { $token, notAllowedTriggered, notFoundTriggered, unauthorizedTriggered } from "@/shared/session";

export type FetchMethods = "GET" | "POST" | "PUT" | "DELETE";
export type BaseResponse = {
  code: number;
  success: boolean;
  error: string;
  data: any;
  meta: {
    page: number;
    count: number;
    pages: number;
  };
};

const BASE_URL = process.env.NODE_ENV === "development" ? "//127.0.0.1:4200" : "";
const abortControllers = new Map<string, AbortController>();

export const baseRequest = async ({
  token,
  url,
  method,
  body,
  headers,
}: {
  token?: string;
  url: string;
  method: FetchMethods;
  body?: BodyInit;
  headers?: HeadersInit;
}) => {
  try {
    abortControllers.get(url)?.abort();
    const controller = new AbortController();
    const { signal } = controller;
    abortControllers.set(url, controller);

    const response = await fetch(`${BASE_URL}${url}`, {
      method,
      credentials: "include",
      referrerPolicy: "origin",
      body,
      headers: { ...headers, "Content-Type": "application/json", Authorization: `Bearer ${token}` },
      signal,
    });

    const responseText = await response.clone().text();

    switch (response.status) {
      case 401:
        unauthorizedTriggered();
        break;
      case 403:
        notAllowedTriggered();
        break;
      case 404:
        notFoundTriggered();
        break;
    }

    const jsonResponse = JSON.parse(responseText);

    return jsonResponse;
  } catch (error) {
    console.log(error);
  } finally {
    abortControllers.delete(url);
  }
};

export const get = attach({
  source: $token,
  mapParams: ({ url, headers }: { url: string; headers?: HeadersInit }, token) => ({ token, url, headers }),
  effect: createEffect(({ token, url, headers }: { token: string; url: string; headers?: HeadersInit }) => {
    return baseRequest({ token, url, method: "GET", headers });
  }),
});

export const post = attach({
  source: $token,
  mapParams: ({ url, body, headers }: { url: string; body: BodyInit; headers?: HeadersInit }, token) => ({
    token,
    url,
    body,
    headers,
  }),
  effect: createEffect(({ token, url, body, headers }: { token: string; url: string; body: BodyInit; headers?: HeadersInit }) => {
    return baseRequest({ token, url, method: "POST", body, headers });
  }),
});

export const put = attach({
  source: $token,
  mapParams: ({ url, body, headers }: { url: string; body: BodyInit; headers?: HeadersInit }, token) => ({
    token,
    url,
    body,
    headers,
  }),
  effect: createEffect(({ token, url, body, headers }: { token: string; url: string; body: BodyInit; headers?: HeadersInit }) => {
    return baseRequest({ token, url, method: "PUT", body, headers });
  }),
});

export const del = attach({
  source: $token,
  mapParams: ({ url, headers }: { url: string; headers?: HeadersInit }, token) => ({
    token,
    url,
    headers,
  }),
  effect: createEffect(({ token, url, headers }: { token: string; url: string; headers?: HeadersInit }) => {
    return baseRequest({ token, url, method: "DELETE", headers });
  }),
});
