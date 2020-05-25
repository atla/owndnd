<style>
.loginText {
    color: greenyellow;
}
</style>

<script>
import {user} from "./stores.js"

</script>

<div class="center-align">
  <h4>Welcome {$user.name}</h4>
  <p>Please <a href class=loginText>log</a> in or <a href class=loginText>signup</a></p>
</div>
