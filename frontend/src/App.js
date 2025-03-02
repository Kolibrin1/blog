import { render, screen, fireEvent } from "@testing-library/react";
import LikeButton from "./components/Blog/LikeButton.js";
import Login from "./components/Auth/Login.js";

// Тест для LikeButton
it("должен увеличивать количество лайков при клике", () => {
  render(<LikeButton />);

  const button = screen.getByRole("button");
  const initialLikes = screen.getByText(/Likes:/i);

  expect(initialLikes).toHaveTextContent("Likes: 0");

  fireEvent.click(button);
  expect(screen.getByText(/Likes:/i)).toHaveTextContent("Likes: 1");
});

// Тест для Login
it("должен вызывать onSubmit при отправке формы", () => {
  const mockSubmit = jest.fn();
  render(<Login onSubmit={mockSubmit} />);

  const emailInput = screen.getByLabelText(/email/i);
  const passwordInput = screen.getByLabelText(/password/i);
  const submitButton = screen.getByRole("button", { name: /login/i });

  fireEvent.change(emailInput, { target: { value: "test@example.com" } });
  fireEvent.change(passwordInput, { target: { value: "password123" } });
  fireEvent.click(submitButton);

  expect(mockSubmit).toHaveBeenCalledWith({
    email: "test@example.com",
    password: "password123",
  });
});
