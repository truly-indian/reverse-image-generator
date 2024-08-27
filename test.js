const OpenAI = require('openai')

const openai = new OpenAI({
	apiKey: "pk-ecSQdYGHiEnfgPwFudtRbkjcpfEqEfAxgPMmHIYlYDHUUFsQ",
	baseURL: "https://api.pawan.krd/v1",
});

(async () => {
    const chatCompletion = await openai.chat.completions.create({
      messages: [
        {"role": "user", "content": "How do I list all files in a directory using Python?"},
    ],
        model: 'pai-001-light',
      });
      console.log(chatCompletion)
      console.log(chatCompletion.choices[0].message.content);
})();

