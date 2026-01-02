export default function RegisterView(){
    return(`
        
         <div class="auth-container">
      <h2>Register</h2>

      <form id="login-form" class="auth-form">
        <input
          type="text"
          name="email"
          placeholder="Email"
          required
        />
          <input
          type="text"
          name="username"
          placeholder="Username"
          required
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          required
        />

        <button type="submit">Login</button>
      </form>

      <p class="auth-footer">
        Already have an account? <a href="/login">Login</a>
      </p>
    </div>
  `)

}