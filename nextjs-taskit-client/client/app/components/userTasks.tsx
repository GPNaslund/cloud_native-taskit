"use client";
import { useEffect, useState } from "react";
import Task from "./task";
import endpointService from "../services/endpointService";
import FlashBox, { FlashBoxProps, FlashBoxType } from "./flashBox";

/**
 * Represents a task.
 */
type Task = {
  id: string;
  title: string;
  details: string;
  is_done: boolean;
};

/**
 * Represents the UserTasks component that displays the users tasks.
 */
export default function UserTasks() {
  // The base path that the client is running on
  const basePath = endpointService.getBasePath();

  // Variable for holding the tasks of the user.
  const [tasks, setTasks] = useState<Task[]>([]);

  // Variable for keeping track of loading state.
  const [loading, setLoading] = useState(true);

  // Variable for holiding the flash box props.
  const [flashBoxProps, setFlashBoxProps] = useState<FlashBoxProps>({
    content: "",
    type: FlashBoxType.Success,
    trigger: 0,
  });

  /**
   * Method for executing functionality once on render.
   */
  useEffect(() => {
    getUserTasks();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  /**
   * Functionality for getting all the tasks of the user.
   */
  async function getUserTasks() {
    try {
      const res = await fetch(endpointService.getTasksUrl(), {
        credentials: "include",
      });

      if (res.ok) {
        const json = await res.json();
        console.log(json);
        setTasks(json);
        setLoading(false);
      } else if (res.status == 401) {
        sessionStorage.setItem("flash-type", "info");
        sessionStorage.setItem("flash-message", "You need to log in");
        window.location.href = basePath;
      } else {
        setFlashBoxProps({
          content: "Something went wrong, try again later..",
          type: FlashBoxType.Danger,
          trigger: flashBoxProps.trigger + 1,
        });
      }
    } catch (error) {
      console.error(error);
      setFlashBoxProps({
        content: "Something went wrong, try again later..",
        type: FlashBoxType.Danger,
        trigger: flashBoxProps.trigger + 1,
      });
    }
  }

  /**
   * Functionality for handling the deletion of a task.
   */
  function handleTaskDelete(statusCode: number, taskId: string) {
    if (statusCode == 204) {
      // Remove deleted task
      setTasks(tasks.filter((task) => task.id !== taskId));
      setFlashBoxProps((prev) => ({
        content: "Deleted task successfully",
        type: FlashBoxType.Success,
        trigger: prev.trigger + 1,
      }));
    } else if (statusCode == 401) {
      sessionStorage.setItem("flash-type", "info");
      sessionStorage.setItem("flash-message", "You need to log in!");
      window.location.href = basePath;
    } else {
      setFlashBoxProps((prev) => ({
        content: "Something went wrong, try again later..",
        type: FlashBoxType.Danger,
        trigger: prev.trigger + 1,
      }));
    }
  }

  function handleTaskUpdate(statusCode: number) {
    if (statusCode == 200) {
      setFlashBoxProps((prev) => ({
        content: "Task updated successfully",
        type: FlashBoxType.Success,
        trigger: prev.trigger + 1,
      }));
    } else if (statusCode == 401) {
      sessionStorage.setItem("flash-type", "info");
      sessionStorage.setItem("flash-message", "You need to log in!");
      window.location.href = basePath;
    } else {
      setFlashBoxProps((prev) => ({
        content: "Something went wrong, try again later..",
        type: FlashBoxType.Danger,
        trigger: prev.trigger + 1,
      }));
      window.location.href = basePath;
    }
  }

  if (loading) {
    return <div aria-busy="true"></div>;
  }

  if (tasks.length === 0) {
    return <div className="container">No tasks found. Create some!</div>;
  }

  return (
    <div className="container">
      <FlashBox {...flashBoxProps} />
      {tasks.map((task) => (
        <Task
          key={task.id}
          id={task.id}
          title={task.title}
          details={task.details}
          isDone={task.is_done}
          onDelete={(statusCode: number) =>
            handleTaskDelete(statusCode, task.id)
          }
          onUpdate={(statusCode: number) => handleTaskUpdate(statusCode)}
        />
      ))}
    </div>
  );
}
