export type FetchMethods = "GET" | "POST" | "PUT";

const BASE_URL = process.env.NODE_ENV === "development" ? "//127.0.0.1:4200" : "";

export const baseRequest = async ({
  url,
  method,
  body,
  headers,
}: {
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
      headers: { ...headers },
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

export const get = async ({ url, headers }: { url: string; headers?: HeadersInit }) => {
  return baseRequest({ url, method: "GET", headers });
};

export const post = async ({ url, body, headers }: { url: string; body: BodyInit; headers?: HeadersInit }) => {
  return baseRequest({ url, method: "POST", body, headers });
};

export const put = async ({ url, body, headers }: { url: string; body: BodyInit; headers?: HeadersInit }) => {
  return baseRequest({ url, method: "PUT", body, headers });
};
