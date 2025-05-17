import Link from "next/link";
import Script from "next/script";

export default function Home() {
  return (
    <>
      <header className="header">
        <nav>
          <div className="logo-header"></div>
          <div className="menu-header">
            <a>Создать пост</a>
          </div>
          <div className="profile-header">
            <a>Профиль</a>
          </div>
        </nav>
      </header>
      <Script src="/js/scroll.js" />
    </>
  );
}
