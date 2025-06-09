import React, { useEffect, useState } from 'react'
import axios from 'axios'

function App() {
  const [posts, setPosts] = useState([])
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  useEffect(() => {
    axios.get('/api/posts').then(res => setPosts(res.data))
  }, [])

  const handleCreate = async () => {
    await axios.post('/api/posts', { title, content })
    const res = await axios.get('/api/posts')
    setPosts(res.data)
    setTitle('')
    setContent('')
  }

  const handleDelete = async (id) => {
    await axios.delete(`/api/posts/${id}`)
    setPosts(posts.filter(p => p.id !== id))
  }

  return (
    <div style={{ padding: 20 }}>
      <h1>📝 My Blog</h1>
      <input
        value={title}
        onChange={e => setTitle(e.target.value)}
        placeholder="Title"
      />
      <br />
      <textarea
        value={content}
        onChange={e => setContent(e.target.value)}
        placeholder="Content"
      />
      <br />
      <button onClick={handleCreate}>Create Post</button>

      <hr />

      <ul>
        {posts.map(post => (
          <li key={post.id}>
            <strong>{post.title}</strong>: {post.content}
            <button onClick={() => handleDelete(post.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App
