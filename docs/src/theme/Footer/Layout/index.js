import React from "react";
import clsx from "clsx";
export default function FooterLayout({ style, links, logo, copyright }) {
  return (
    <div
      className={clsx("footer", {
        "footer--dark": style === "dark",
      })}
    >
      <div className="flex row container container-fluid">
        {copyright}
        {links}
      </div>
    </div>
  );
}
