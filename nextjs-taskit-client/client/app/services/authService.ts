import endpointService from "./endpointService";

/**
 * Represents service class for checking if user is authenticated.
 */
class AuthService {
  private lastRefresh = Date.now();
  private readonly REFRESH_THRESHOLD = 15 * 60 * 1000; // 15min
  private authUrl = endpointService.getAuthUrl();

  /**
   * Method for checking if user is authenticated.
   */
  async isAuthenticated(): Promise<boolean> {
    try {
      const res = await fetch(this.authUrl, {
        credentials: "include",
      });

      if (!res.ok) {
        return false;
      }

      const timeSinceRefresh = Date.now() - this.lastRefresh;
      if (timeSinceRefresh > this.REFRESH_THRESHOLD) {
        await this.refreshSession();
      }

      return true;
    } catch (error) {
      console.error("Auth check failed:", error);
      return false;
    }
  }

  private async refreshSession(): Promise<void> {
    try {
      const res = await fetch(`${this.authUrl}/refresh`, {
        method: "POST",
        credentials: "include",
      });
      if (res.ok) {
        this.lastRefresh = Date.now();
        window.location.reload();
      }
    } catch (error) {
      console.error("Session refresh failed: ", error);
    }
  }
}

const authService = new AuthService();
export default authService;
