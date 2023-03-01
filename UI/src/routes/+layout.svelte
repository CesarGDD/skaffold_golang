<script lang="ts">
	import type { Post } from './+layout';
  let username = '';
  let title = '';
  let content = '';
  let posts: Post[] = [];
  let post2: Post = {
	id: 0,
	title: '', 
	content: '', 
	username: ''
};

  async function createUser() {
    try {
      const response = await fetch('http://localhost:3000/user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username })
      });

      if (!response.ok) {
        throw new Error(response.statusText);
      }
	  const user = await response.json();
	  post2.id = Math.floor(Math.random() * 100);
	  post2.username = user.username;
      console.log('User created successfully!');
    } catch (error) {
      console.error('Error creating user:', error);
    }
  }

  async function createPost() {
    try {
      const response = await fetch('http://localhost:3000/post', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ title, content })
      });

      if (!response.ok) {
        throw new Error(response.statusText);
      }

      const post = await response.json();
	  const newPost = {
      id: Math.floor(Math.random() * 100),
      title: post.title,
      content: post.content,
      username: post2.username
    };
    posts = [...posts, newPost];
      console.log('Post created successfully!');
    } catch (error) {
      console.error('Error creating post:', error);
    }
  }

  const createPostAndUser  = async () => {
	createUser();
	createPost();
  };
</script>

<style>
	.container {
	  display: flex;
	  flex-direction: column;
	  align-items: center;
	  justify-content: center;
	  height: 100vh;
	}
  
	label {
	  color: #414141;
	  font-weight: bold;
	  margin-top: 8px;
	  margin-bottom: 4px;
	}
  
	input,
	textarea {
	  background-color: #36393f;
	  border: none;
	  border-radius: 4px;
	  color: #dcddde;
	  padding: 8px;
	  width: 100%;
	  margin-bottom: 16px;
	  box-sizing: border-box;
	}
  
	textarea {
	  height: 150px;
	}
  
	button {
	  background-color: #7289da;
	  border: none;
	  border-radius: 4px;
	  color: #fff;
	  font-weight: bold;
	  padding: 8px 16px;
	  cursor: pointer;
	  transition: background-color 0.2s ease-in-out;
	  margin-top: 16px;
	}
  
	button:hover {
	  background-color: #677bc4;
	}
  </style>
  
<div class="container">
	<div>
		<label for="username">Username:</label>
		<input type="text" id="username" bind:value={username} />
</div>

<div>
  <label for="title">Title:</label>
  <input type="text" id="title" bind:value={title} />
</div>

<div>
	<label for="content">Content:</label>
	<textarea id="content" bind:value={content}></textarea>
</div>

<button on:click={createPostAndUser}>Create Post</button>

<h2>Posts:</h2>

{#if posts.length === 0}
  <p>No posts yet</p>
{:else}
  {#each posts as post}
    <div>
      <h3>{post.title}</h3>
      <p>{post.content}</p>
      <p>By {post.username}</p>
    </div>
  {/each}
{/if}
</div>