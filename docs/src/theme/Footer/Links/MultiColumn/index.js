import React from "react";
import LinkItem from "@theme/Footer/LinkItem";
function RowLinkItem({ item }) {
  return item.html ? (
    <li
      className="footer__item"
      // Developer provided the HTML, so assume it's safe.
      // eslint-disable-next-line react/no-danger
      dangerouslySetInnerHTML={{ __html: item.html }}
    />
  ) : (
    <li
      key={item.href ?? item.to}
      className="footer__item"
      style={{
        display: "inline-block",
        marginRight: "1rem",
      }}
    >
      <LinkItem item={item} />
    </li>
  );
}
function Column({ column }) {
  return (
    <div>
      <ul className="footer__items clean-list">
        {column.items.map((item, i) => (
          <RowLinkItem key={i} item={item} />
        ))}
      </ul>
    </div>
  );
}
export default function FooterLinksMultiColumn({ columns }) {
  return (
    <div className="footer__links">
      {columns.map((column, i) => (
        <Column key={i} column={column} />
      ))}
    </div>
  );
}
