import Link from "next/link";

/**
 * Represents the nav bar props.
 */
type NavBarProps = {
  isLoggedIn: boolean;
};

/**
 * Represents the NavBar component with links for navigation.
 *
 * @param isLoggedIn - Boolean value indicating if user is logged in or not.
 */
export default function NavBar({ isLoggedIn }: NavBarProps) {
  if (isLoggedIn) {
    return (
      <nav className="container">
        <ul>
          <li>
            <strong>TASKIT</strong>
          </li>
        </ul>
        <ul>
          <li>
            <Link href="/">Home</Link>
          </li>
          <li>
            <Link href="/add-task">Add Task</Link>
          </li>
          <li>
            <Link href="/logout">Logout</Link>
          </li>
        </ul>
      </nav>
    );
  } else {
    return (
      <nav className="container">
        <ul>
          <li>
            <strong>TASKIT</strong>
          </li>
        </ul>
        <ul>
          <li>
            <Link href="/">Login</Link>
          </li>
          <li>
            <Link href="/register">Register</Link>
          </li>
        </ul>
      </nav>
    );
  }
}
