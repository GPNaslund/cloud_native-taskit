"use client";
import { useEffect, useState } from "react";
import style from "./flashBox.module.css";

/**
 * Represents the type of flashbox to display.
 */
export enum FlashBoxType {
  Success = "flash-success",
  Danger = "flash-danger",
  Info = "flash-info",
}

/**
 * Represents the props for the flash box.
 */
export type FlashBoxProps = {
  content: string;
  type: FlashBoxType;
  trigger: number;
};

/**
 * Represents the flashbox component.
 *
 * @param content - The text content to display.
 * @param type - The type of FlashBox to display.
 * @param trigger - An integer repsenting the number of flashbox to display.
 */
export default function FlashBox({ content, type, trigger }: FlashBoxProps) {
  // Variable for keeping track if the flashbox is currently visible.
  const [isVisible, setIsVisible] = useState(false);

  // Variable for keeping track of where in the animation cycle the flash box is.
  const [animationClass, setAnimationClass] = useState("");

  // The current content to display.
  const [currentContent, setCurrentContent] = useState(content);

  /**
   * Method for executing functionality on each change to trigger
   * and content variables.
   */
  useEffect(() => {
    // Display flashbox if trigger is larger than 0
    if (trigger > 0) {
      setCurrentContent(content);

      setIsVisible(true);
      setAnimationClass(style["fade-in"]);

      const hideSequence = () => {
        setAnimationClass(style["fade-out"]);
        setTimeout(() => {
          setIsVisible(false);
          setAnimationClass("");
        }, 500);
      };

      const timer = setTimeout(hideSequence, 3000);

      return () => clearTimeout(timer);
    }
  }, [trigger, content]);

  // Hide flashbox if not visible.
  if (!isVisible) return null;

  return (
    <div className={`${style.root} ${style[type]} ${animationClass}`}>
      <p className={style["flash-content"]}>{currentContent}</p>
    </div>
  );
}
