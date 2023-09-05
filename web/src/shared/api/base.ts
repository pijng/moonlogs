import { attach, createEffect } from "effector";
import { $token } from "../auth";

export type FetchMethods = "GET" | "POST" | "PUT";

const BASE_URL = process.env.NODE_ENV === "development" ? "//127.0.0.1:4200" : "";

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
    const response = await fetch(`${BASE_URL}${url}`, {
      method,
      referrerPolicy: "origin-when-cross-origin",
      body,
      headers: { ...headers, Authorization: `Bearer ${token}` },
    });

    const responseText = await response.clone().text();

    if (!response.ok) {
      throw new Error(responseText);
    }

    if (response.status === 204) return null;

    const jsonResponse = JSON.parse(responseText);

    console.log(jsonResponse);
    return jsonResponse;
  } catch (error) {
    console.log(error);
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
