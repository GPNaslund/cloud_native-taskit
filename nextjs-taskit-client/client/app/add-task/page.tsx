"use client";

import { useEffect, useState } from "react";
import NavBar from "../components/navbar";
import endpointService from "../services/endpointService";
import { useRouter } from "next/navigation";
import authService from "../services/authService";

/**
 * Represents the add task page.
 */
export default function AddTaskPage() {
  // Variable for keeping track if auth status has been checked.
  const [authIsChecked, setAuthIsChecked] = useState(false);

  // Variable for keeping track of the title value of the task to add.
  const [titleValue, setTitleValue] = useState("");

  // Variable for keeping track of the details value of the task to add.
  const [detailsValue, setDetailsValue] = useState("");

  // Router instance for ability to navigate.
  const router = useRouter();

  /**
   * Method for handling the add task action.
   *
   * @param e - The react form event.
   */
  async function handleSubmit(e: React.FormEvent<HTMLInputElement>) {
    e.preventDefault();
    const addTaskUrl = endpointService.getAddTaskUrl();

    // Send POST request to backend for adding new task.
    const res = await fetch(addTaskUrl, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: titleValue,
        details: detailsValue,
        is_done: false,
      }),
    });

    if (res.ok) {
      sessionStorage.setItem("flash-type", "success");
      sessionStorage.setItem("flash-message", "Added task successfully!");
      router.push("/");
    } else if (res.status == 401) {
      sessionStorage.setItem("flash-type", "info");
      sessionStorage.setItem(
        "flash-message",
        "You are not logged in, please log in!",
      );
      router.push("/");
    } else {
      sessionStorage.setItem("flash-type", "danger");
      sessionStorage.setItem(
        "flash-message",
        "Something went wrong, try again later",
      );
      router.push("/");
    }
  }

  /**
   * Method for running auth check.
   */
  useEffect(() => {
    checkAuth();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  /*
   * Method for checking authentication status. Redirects to start page
   * if not authenticated.
   */
  async function checkAuth() {
    const isAuthenticated = await authService.isAuthenticated();
    if (!isAuthenticated) {
      router.push("/");
    } else {
      setAuthIsChecked(true);
    }
  }

  if (!authIsChecked) {
    return (
      <div>
        <NavBar isLoggedIn={true} />
        <div id="loading-div" aria-busy="true"></div>
      </div>
    );
  } else {
    return (
      <div className="container">
        <NavBar isLoggedIn={true} />
        <form>
          <fieldset>
            <label>
              Title
              <input
                name="title"
                placeholder="Enter task title"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  setTitleValue(e.currentTarget.value);
                }}
              />
            </label>
            <label>
              Details
              <input
                name="email"
                placeholder="Enter details"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  setDetailsValue(e.currentTarget.value);
                }}
              />
            </label>
          </fieldset>

          <input
            type="submit"
            value="Add task"
            className="outline"
            onClick={handleSubmit}
          />
        </form>
      </div>
    );
  }
}
