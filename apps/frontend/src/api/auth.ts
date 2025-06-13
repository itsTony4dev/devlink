export const loginUser = async (email: string, password: string) => {
  const response = await fetch('http://localhost:8080/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  })

  if (!response.ok) {
    const errorData = await response.json()
    throw new Error(errorData.message || 'Login failed')
  }

  return response.json()
}

export const registerUser = async (username: string, email: string, password: string) => {
  const response = await fetch('http://localhost:8080/users/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, email, password }),
  })

  if (!response.ok) {
    const errorData = await response.json()
    throw new Error(errorData.message || 'Registration failed')
  }

  return response.json()
} 