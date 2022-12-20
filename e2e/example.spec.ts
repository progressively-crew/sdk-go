import { test, expect } from "@playwright/test";

const updateFlagStatus = (
  envId: string,
  flagId: string,
  status: string,
  token: string
) => {
  return fetch(`http://localhost:4000/environments/${envId}/flags/${flagId}`, {
    method: "PUT",
    body: JSON.stringify({
      status,
    }),
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
};

const changeFlagStatus = async (
  envId: string,
  flagId: string,
  status: string
) => {
  const response = await fetch("http://localhost:4000/auth/login", {
    method: "POST",
    body: JSON.stringify({
      username: "marvin.frachet@something.com",
      password: "password",
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const { access_token } = await response.json();
  return updateFlagStatus(envId, flagId, status, access_token);
};

test("checks the not activate variant of the flag", async ({ page }) => {
  await page.goto("/");

  await expect(page.getByText("New variant")).toBeTruthy();

  await changeFlagStatus("1", "1", "NOT_ACTIVATED");

  await page.goto("/");
  await expect(page.getByText("Old variant")).toBeTruthy();
});
