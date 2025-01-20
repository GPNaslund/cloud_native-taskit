/**
 * Service class for returning API endpoint urls.
 */
class EndpointService {
  // The base api that is read by env.
  baseUrl: string;

  // The base path that the client is running on
  basePath: string;

  /**
   * Creates a new instance of EndpointService.
   */
  constructor() {
    // Reads the env
    const url = process.env.NEXT_PUBLIC_API_URL;
    if (url == undefined) {
      throw new Error("NEXT_PUBLIC_API_URL must be set");
    }
    this.baseUrl = url;

    const basePath = process.env.NEXT_PUBLIC_BASE_PATH;
    if (basePath == undefined) {
      throw new Error("NEXT_PUBLIC_BASE_PATH must be set");
    }
    this.basePath = basePath;
  }

  /**
   * Method for getting the login api endpoint url.
   */
  getLoginUrl(): string {
    return `${this.baseUrl}/v1/users/login`;
  }

  /**
   * Method for getting the register api endpoint url.
   */
  getRegisterUrl(): string {
    return `${this.baseUrl}/v1/users`;
  }

  /**
   * Method for getting the task api endpoint url.
   */
  getTaskUrl(taskId: string): string {
    return `${this.baseUrl}/v1/users/me/tasks/${taskId}`;
  }

  /**
   * Method for getting the tasks api endpoint url.
   */
  getTasksUrl(): string {
    return `${this.baseUrl}/v1/users/me/tasks`;
  }

  /**
   * Method for getting the add task api endpoint url.
   */
  getAddTaskUrl() {
    return `${this.baseUrl}/v1/users/me/tasks`;
  }

  /**
   * Method for getting the logout api endpoint url.
   */
  getLogoutUrl() {
    return `${this.baseUrl}/v1/users/logout`;
  }

  /**
   * Method for getting the auth api endpoint url.
   */
  getAuthUrl() {
    return `${this.baseUrl}/v1/auth/session`;
  }

  /**
   * Method for getting the base path.
   */
  getBasePath() {
    return this.basePath;
  }
}

const endpointService = new EndpointService();
export default endpointService;
