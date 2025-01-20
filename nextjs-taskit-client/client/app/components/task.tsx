"use client";
import style from "./task.module.css";
import React, { useState } from "react";
import endpointService from "../services/endpointService";

/**
 * Represents the props of a task.
 */
export type TaskProps = {
  id: string;
  title: string;
  details: string;
  isDone: boolean;
  onDelete: (statusCode: number) => void;
  onUpdate: (statusCode: number) => void;
};

/**
 * Represents a task component.
 *
 * @param id - The id of the task.
 * @param title - The title of the task.
 * @param details - The details of the task.
 * @param isDone - Boolean value indicating if the task is done.
 * @param onDelete - Callback function that gets called when task is deleted.
 */
export default function Task({
  id,
  title,
  details,
  isDone,
  onDelete,
  onUpdate,
}: TaskProps) {
  // Variable for keeping track of the IsDone checkbox state.
  const [isChecked, setIsChecked] = useState(isDone);

  // Variable for keeping track if the task is in edit mode.
  const [isInEditMode, setIsInEditMode] = useState(false);

  // Variable for keeping track of the title value of the task.
  const [titleValue, setTitleValue] = useState(title);

  // Variable for keeping track of the details value of the task.
  const [detailsValue, setDetailsValue] = useState(details);

  // Variable for keeping track if deletion is being performed.
  const [isDeleting, setIsDeleting] = useState(false);

  /**
   * Method for handling the state of the IsDone checkbox.
   */
  function handleCheckboxChange() {
    setIsChecked(!isChecked);
  }

  /**
   * Method for handling the save changes of task logic.
   */
  async function handleSave() {
    try {
      const res = await fetch(endpointService.getTaskUrl(id), {
        method: "PUT",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: id,
          title: titleValue,
          details: detailsValue,
          is_done: isChecked,
        }),
      });

      setIsInEditMode(false);
      onUpdate(res.status);
    } catch (error) {
      console.error("Error updating task:", error);
    }
  }

  /**
   * Method for toggling the edit mode of the task.
   */
  function toggleEditMode() {
    setIsInEditMode(!isInEditMode);
    if (isInEditMode) {
      setIsChecked(isDone);
      setTitleValue(title);
      setDetailsValue(details);
    }
  }

  /**
   * Method for handling the deletion of the task.
   */
  async function handleDelete() {
    if (isDeleting) return;
    try {
      setIsDeleting(true);
      const res = await fetch(endpointService.getTaskUrl(id), {
        method: "DELETE",
        credentials: "include",
      });
      onDelete(res.status);
      setIsDeleting(false);
      setIsInEditMode(false);
    } catch (error) {
      console.error("Error deleting task:", error);
      setIsDeleting(false);
      setIsInEditMode(false);
    }
  }

  return (
    <article className={style.root}>
      <header className={style.header}>
        <input
          type="text"
          value={titleValue}
          className={style.input}
          disabled={!isInEditMode}
          onChange={(e) => {
            setTitleValue(e.currentTarget.value);
          }}
        />
        <label>
          <input
            type="checkbox"
            name="isDone"
            required
            checked={isChecked}
            onChange={handleCheckboxChange}
            disabled={!isInEditMode}
          />
          Done
        </label>
      </header>
      <input
        type="text"
        required
        value={detailsValue}
        className={style.input}
        disabled={!isInEditMode}
        onChange={(e) => setDetailsValue(e.currentTarget.value)}
      />
      <footer className={style.footer}>
        <div className={isInEditMode ? "hidden" : ""}>
          <button className="outline" onClick={toggleEditMode}>
            Edit task
          </button>
        </div>
        <div className={isInEditMode ? style["cancelSaveContainer"] : "hidden"}>
          <button
            className="outline"
            onClick={handleDelete}
            disabled={isDeleting}
          >
            {isDeleting ? "Deleting..." : "Delete"}
          </button>
          <button className="outline contrast" onClick={toggleEditMode}>
            Cancel
          </button>
          <button
            className="outline"
            onClick={handleSave}
            disabled={!titleValue || !detailsValue}
          >
            Save
          </button>
        </div>
      </footer>
    </article>
  );
}
