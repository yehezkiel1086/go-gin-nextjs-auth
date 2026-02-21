export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface RegisterResponse {
  message?: string;
  error?: string;
}

export interface Job {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  title: string;
  description: string;
  company: string;
  location: string;
  salary: number;
}

export interface GetJobsResponse {
  jobs?: Job[];
  error?: string;
}

const API_BASE_URL = 'http://127.0.0.1:8080/api/v1';

export async function register(userData: RegisterRequest): Promise<RegisterResponse> {
  try {
    const response = await fetch(`${API_BASE_URL}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Register error:', error);
    throw error;
  }
}

export async function getJobs(): Promise<GetJobsResponse> {
  try {
    const response = await fetch(`${API_BASE_URL}/jobs`, {
      method: 'GET',
      credentials: "include",
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return { jobs: data };
  } catch (error) {
    console.error('Get jobs error:', error);
    throw error;
  }
}
