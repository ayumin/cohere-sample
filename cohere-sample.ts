import { CohereClientV2 } from "cohere-ai";
import * as dotenv from "dotenv";

dotenv.config();

const co = new CohereClientV2({
  token: process.env.COHERE_API_KEY,
});

const documents = [
  {
    data: {
      title: "会社概要",
      text: "私たちはAIを活用したソフトウェア開発を行う企業です。2020年に設立されました。",
    },
  },
  {
    data: {
      title: "サービス",
      text: "主なサービスはチャットボット開発、データ分析、AI導入支援です。",
    },
  },
];

const response = await co.chat({
  model: "command-a-03-2025",
  messages: [
    {
      role: "system",
      content:
        "与えられた文書に基づいて回答してください。不明な点は推測せず不明と答えてください。",
    },
    {
      role: "user",
      content: "この会社はどんなサービスを提供していますか？",
    },
  ],
  documents,
});

// 本文表示
if (response.message?.content) {
  const textContent = response.message.content.find((c) => c.type === "text");
  if (textContent && textContent.type === "text") {
    console.log(textContent.text);
  }
}

// citation 表示
if (response.message?.citations) {
  for (const citation of response.message.citations) {
    console.log(citation);
  }
}
