import { BaseResponse, post } from "./base";

export interface InsightResponse extends BaseResponse {
  data: string;
}

type ContentPayload = {
  prompt: string;
};

export const generateContent = (prompt: string): Promise<InsightResponse> => {
  const body: ContentPayload = { prompt };
  return post({ url: "/api/insights/generate_content", body: JSON.stringify(body) });
};
