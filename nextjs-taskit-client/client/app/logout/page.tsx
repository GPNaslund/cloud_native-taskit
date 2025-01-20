"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";
import endpointService from "../services/endpointService";

/**
 * Represents the logout page.
 */
export default function LogoutPage() {
  // Router instance for client side routing.
  const router = useRouter();

  /**
   * Method for executing functionality
   */
  useEffect(() => {
    async function handleLogout() {
      try {
        const res = await fetch(endpointService.getLogoutUrl(), {
          method: "DELETE",
          credentials: "include",
        });

        if (res.ok) {
          sessionStorage.setItem("flash-type", "success");
          sessionStorage.setItem("flash-message", "Logged out successfully");
          router.push("/");
        } else {
          console.log("Logout failed");
          sessionStorage.setItem("flash-type", "info");
          sessionStorage.setItem(
            "flash-message",
            "Logout failed but session forced removed",
          );
          router.push("/");
        }
      } catch (error) {
        console.error("Logout error:", error);
        sessionStorage.setItem("flash-type", "danger");
        sessionStorage.setItem(
          "flash-message",
          "Failed to log out. Please try again.",
        );
        router.push("/");
      }
    }

    handleLogout();
  }, [router]);

  return <div aria-busy="true">Logging out...</div>;
}
