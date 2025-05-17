import Script from "next/script";

export default function Form() {
  return (
    <>
      <div className="container" id="container">
        <div className="form-container sign-up">
          <form>
            <h1>Create Account</h1>
            <input type="text" placeholder="surname"></input>
            <input type="text" placeholder="name"></input>
            <label htmlFor="file-upload" className="file-upload-button">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="20"
                height="20"
                fill="currentColor"
                className="bi bi-upload"
                viewBox="0 0 16 16"
              >
                <path d="M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5z" />
                <path d="M7.646 1.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708L8.5 2.707V11.5a.5.5 0 0 1-1 0V2.707L5.354 4.854a.5.5 0 1 1-.708-.708l3-3z" />
              </svg>
              Загрузить аватар
            </label>
            <input
              type="file"
              accept=".jpg, .jpeg, .png"
              id="file-upload"
            ></input>
            <input type="email" placeholder="email" required></input>
            <input type="text" placeholder="login" required></input>
            <input type="password" placeholder="password" required></input>
            <input
              type="password"
              placeholder="repeat password"
              required
            ></input>
            <button>Sign Up</button>
          </form>
        </div>
        <div className="form-container sign-in">
          <form>
            <h1>Sign In</h1>
            <input type="email" placeholder="email" required></input>
            <input type="text" placeholder="login" required></input>
            <input type="password" placeholder="password" required></input>
            <button>Sign In</button>
          </form>
        </div>
        <div className="toggle-container">
          <div className="toggle">
            <div className="toggle-panel toggle-left">
              <h1>Welcom Back!</h1>
              <p>Enter your personal details to use all of site features</p>
              <button className="hiden" id="login">
                Sign In
              </button>
            </div>
            <div className="toggle-panel toggle-right">
              <h1>Hello, Friend!</h1>
              <p>
                Register with your personal details to use all of site features
              </p>
              <button className="hiden" id="register">
                Sign Up
              </button>
            </div>
          </div>
        </div>
      </div>
      <Script src="/js/script.js" />
    </>
  );
}
