"use client";
import { useState } from "react";
import FlashBox, { FlashBoxProps, FlashBoxType } from "./flashBox";
import endpointService from "../services/endpointService";

/**
 * Represents the login page.
 * @param onLoginSuccess - Callback function that gets called on successful login.
 */
export default function LoginPage() {
  // The API url to send login request to.
  const loginUrl = endpointService.getLoginUrl();

  // The base path of the client
  const basePath = endpointService.getBasePath();

  // Reference to the username input element.
  const [usernameValue, setUsernameValue] = useState("");

  // Reference to the password input element.
  const [passwordValue, setPasswordValue] = useState("");

  // Variables to keep track of login state.
  const [isLoggingIn, setIsLoggingIn] = useState(false);

  // Variables for keeping track of flash box props.
  const [flashProps, setFlashProps] = useState<FlashBoxProps>({
    content: "",
    type: FlashBoxType.Info,
    trigger: 0,
  });

  /**
   * Method for handling login submission.
   * @param event - The mouse event from the HTML input element.
   */
  async function handleClick(event: React.MouseEvent<HTMLInputElement>) {
    event.preventDefault();

    // Do nothing if log in is in progress.
    if (isLoggingIn) return;

    setIsLoggingIn(true);

    // Make POST request to backend.
    const res = await fetch(loginUrl!, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: usernameValue,
        password: usernameValue,
      }),
      credentials: "include",
    });

    if (res.ok) {
      sessionStorage.setItem("flash-type", "success");
      sessionStorage.setItem("flash-message", "Logged in successfully!");

      window.location.href = basePath;
    } else if (res.status == 400) {
      setFlashProps({
        content: "Invalid credentials, try again",
        type: FlashBoxType.Danger,
        trigger: 1,
      });
      setIsLoggingIn(false);
      window.location.href = basePath;
    } else {
      setFlashProps({
        content: "Something went wrong, try again later",
        type: FlashBoxType.Danger,
        trigger: 1,
      });
      setIsLoggingIn(false);
      window.location.href = basePath;
    }
  }

  return (
    <div className="container">
      <form>
        <FlashBox {...flashProps} />
        <fieldset>
          <label>
            Username
            <input
              name="username"
              placeholder="Username"
              required
              disabled={isLoggingIn}
              onChange={(e) => setUsernameValue(e.currentTarget.value)}
            />
          </label>
          <label>
            Password
            <input
              type="password"
              name="password"
              placeholder="Password"
              required
              disabled={isLoggingIn}
              onChange={(e) => setPasswordValue(e.currentTarget.value)}
            />
          </label>
        </fieldset>
        <input
          type="submit"
          value={isLoggingIn ? "Logging in..." : "Log in"}
          onClick={handleClick}
          disabled={isLoggingIn || usernameValue == "" || passwordValue == ""}
        />
      </form>
    </div>
  );
}
