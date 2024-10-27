import React, { useState } from 'react';
import { Send } from 'lucide-react';

const ChatBot = () => {
  const [messages, setMessages] = useState([]);
  const [inputText, setInputText] = useState('');
  const [mode, setMode] = useState('query'); // query or reply

  const sendMessage = async () => {
    if (!inputText.trim()) return;

    const newMessage = {
      mode: mode,
      content: inputText,
      isUser: true,
    };

    setMessages(prev => [...prev, newMessage]);
    setInputText('');

    try {
      const response = await fetch('http://localhost:8080/message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          mode: mode,
          content: inputText,
        }),
      });

      const data = await response.json();
      
      if (response.ok) {
        setMessages(prev => [...prev, {
          mode: data.mode,
          content: data.content,
          isUser: false,
        }]);
      } else {
        console.error('Error:', data.error);
      }
    } catch (error) {
      console.error('Failed to send message:', error);
    }
  };

  return (
    <div className="max-w-2xl mx-auto p-4 h-screen flex flex-col">
      <div className="mb-4 flex justify-between items-center">
        <h1 className="text-2xl font-bold">チャットボット</h1>
        <select 
          value={mode}
          onChange={(e) => setMode(e.target.value)}
          className="border rounded p-2"
        >
          <option value="query">質問モード</option>
          <option value="reply">返答モード</option>
        </select>
      </div>

      <div className="flex-1 border rounded-lg p-4 mb-4 overflow-y-auto bg-gray-50">
        {messages.map((message, index) => (
          <div
            key={index}
            className={`mb-2 ${
              message.isUser ? 'text-right' : 'text-left'
            }`}
          >
            <div
              className={`inline-block rounded-lg px-4 py-2 max-w-[70%] ${
                message.isUser
                  ? 'bg-blue-500 text-white'
                  : 'bg-white border'
              }`}
            >
              {message.content}
            </div>
          </div>
        ))}
      </div>

      <div className="flex gap-2">
        <input
          type="text"
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          onKeyPress={(e) => e.key === 'Enter' && sendMessage()}
          placeholder={mode === 'query' ? '質問を入力してください' : '返答を登録してください'}
          className="flex-1 border rounded-lg p-3"
        />
        <button
          onClick={sendMessage}
          className="bg-blue-500 text-white rounded-lg px-4 py-2 flex items-center gap-2 hover:bg-blue-600"
        >
          <Send size={20} />
          送信
        </button>
      </div>
    </div>
  );
};

export default ChatBot;
