"use client";

import { useState } from "react";
import NavBar from "../components/navbar";
import { useRouter } from "next/navigation";
import FlashBox, { FlashBoxProps, FlashBoxType } from "../components/flashBox";
import endpointService from "../services/endpointService";

/**
 * Represents the register user page.
 */
export default function RegisterPage() {
  // Get the API endpoint url
  const registerUrl = endpointService.getRegisterUrl();

  // Variable for holding the user name value.
  const [usernameValue, setUsernameValue] = useState("");

  // Variable for holding the password value.
  const [passwordValue, setPasswordValue] = useState("");

  // Router instance for client side navigation.
  const router = useRouter();

  // Variable for holding the flash props.
  const [flashProps, setFlashProps] = useState<FlashBoxProps>({
    content: "",
    type: FlashBoxType.Info,
    trigger: 0,
  });

  /*
   * Function for handling the submission click.
   */
  async function handleClick(event: React.MouseEvent<HTMLInputElement>) {
    event.preventDefault();

    const res = await fetch(registerUrl!, {
      method: "post",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: usernameValue,
        password: passwordValue,
      }),
    });

    if (res.ok) {
      sessionStorage.setItem("flash-type", "success");
      sessionStorage.setItem("flash-message", "Registered successfully!");
      router.push("/");
    } else {
      if (res.status == 400) {
        setFlashProps((prev) => ({
          content: "Username allready taken, try another one",
          type: FlashBoxType.Danger,
          trigger: prev.trigger + 1,
        }));
      } else {
        setFlashProps((prev) => ({
          content: "Something went wrong.. try again later",
          type: FlashBoxType.Danger,
          trigger: prev.trigger + 1,
        }));
      }
    }
  }

  return (
    <div className="container">
      <NavBar isLoggedIn={false} />
      <form>
        <FlashBox {...flashProps} />
        <fieldset>
          <label>
            Username
            <input
              name="username"
              placeholder="Username"
              required
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
              onChange={(e) => setPasswordValue(e.currentTarget.value)}
            />
          </label>
        </fieldset>

        <input
          type="submit"
          value="Register"
          className="outline"
          onClick={handleClick}
          disabled={usernameValue == "" || passwordValue == ""}
        />
      </form>
    </div>
  );
}
