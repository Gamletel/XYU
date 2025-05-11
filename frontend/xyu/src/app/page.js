export default function Home() {
  return (
    <>
      <div className="container" id="container">
        <div className="form-container sign-up">
          <form>
            <h1>Create Account</h1>
            <input type="text" placeholder="surname"></input>
            <input type="text" placeholder="name"></input>
            <input type="file" accept=".jpg, .jpeg, .png"></input>
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
    </>
  );
}
