"use client";
import { useEffect, useState } from "react";
import NavBar from "./components/navbar";
import LoginPage from "./components/loginPage";
import UserTasks from "./components/userTasks";
import FlashBox, { FlashBoxProps, FlashBoxType } from "./components/flashBox";
import authService from "./services/authService";

/**
 * Represents the home page.
 */
export default function Home() {
  // Variable for keeping track on if auth check has been done.
  const [authIsChecked, setAuthIsChecked] = useState(false);

  // Variable for keeping track on if user is authenticated.
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  // Variable for keeping track on flash box props.
  const [flashBoxProps, setFlashBoxProps] = useState<FlashBoxProps>({
    content: "",
    type: FlashBoxType.Success,
    trigger: 0,
  });

  /**
   * Method for executing functionality once on render.
   */
  useEffect(() => {
    handleFlashMessage();
    checkAuth();
  }, []);

  /**
   * Method for checking if user is authenticated.
   */
  async function checkAuth() {
    try {
      const isAuthed = await authService.isAuthenticated();
      setIsAuthenticated(isAuthed);
    } catch (error) {
      console.error("Auth check failed:", error);
      setIsAuthenticated(false);
    } finally {
      setAuthIsChecked(true);
    }
  }

  /**
   * Method for handling flash messages from session storage.
   */
  function handleFlashMessage() {
    const flashType = sessionStorage.getItem("flash-type");
    const flashMessage = sessionStorage.getItem("flash-message");

    if (flashType && flashMessage) {
      let flashBoxType = FlashBoxType.Info;
      switch (flashType) {
        case "success":
          flashBoxType = FlashBoxType.Success;
          break;
        case "danger":
          flashBoxType = FlashBoxType.Danger;
          break;
      }

      setFlashBoxProps((prev) => ({
        content: flashMessage,
        type: flashBoxType,
        trigger: prev.trigger + 1,
      }));

      sessionStorage.removeItem("flash-type");
      sessionStorage.removeItem("flash-message");
    }
  }

  if (!authIsChecked) {
    return (
      <div>
        <NavBar isLoggedIn={false} />
        <div id="loading-div" aria-busy="true"></div>
      </div>
    );
  }

  return (
    <div>
      <NavBar isLoggedIn={isAuthenticated} />
      <FlashBox {...flashBoxProps} />
      {isAuthenticated ? <UserTasks /> : <LoginPage />}
    </div>
  );
}
