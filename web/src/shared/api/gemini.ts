import { commonRequest } from "./base";

export type GeminiRequest = {
  contents: GeminiContent[];
};

export type GeminiContentResponse = {
  candidates: GeminiCandidate[];
};

export type GeminiCandidate = {
  content: GeminiContent;
};

export type GeminiContent = {
  parts: GeminiPart[];
  role: "user" | "model";
};

export type GeminiPart = {
  text: string;
};

const GEMINI_URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent";

export const generateContent = ({
  token,
  content,
}: {
  token: string;
  content: GeminiRequest;
}): Promise<GeminiContentResponse | null> => {
  const url = `${GEMINI_URL}?key=${token}`;

  return commonRequest({ url, body: content, method: "POST" });
};
