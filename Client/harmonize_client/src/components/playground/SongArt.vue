<template>
  <div class="song-art-wrapper q-ma-lg" v-bind:style="{ width: width }">
    <q-img
      class="song-art shadow-23"
      v-bind:src="image"
      spinner-color="white"
      v-bind:style="{ width: width }"
    />
    <div class="text-white q-pt-md q-pa-sm text-center text-h5 text-weight-bolder">{{ title }}</div>
    <div class="text-white text-center text-h6">{{ album }} · {{ artist }}</div>
    <div class="text-white text-center text-h6">
        <q-avatar size="28px">
          <img v-bind:src="base_url + user.avatar">
        </q-avatar>
        <span> {{ user.name }}</span>
    </div>
  </div>
</template>

<style lang="sass">
</style>

<script>
import { mapState } from 'vuex';

export default {
  name: "song-art",
  props: ["width", "title", "artist", "album", "image", "user"],
  computed: mapState('rest', {
      base_url: state => state.base_url
  })
};
</script>