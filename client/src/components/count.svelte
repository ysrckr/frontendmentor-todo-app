<script lang="ts">
  export let todosLength: number = 0;
  </script>
  


<div>
    <div class="count">
        <span class="count__number">{todosLength}</span>
        <span class="count__text">items left</span>
    </div>
  
    <div class="clear">
        <button class="clear__button">Clear Completed</button>
    </div>
</div>