<template>
    <div>
        <div style="width: 100%">
            <APlayer :audio="audio"></APlayer>
            <div>
<!--                管理歌曲-->
            </div>
        </div>
    </div>
</template>

<script>
    import { APlayer } from '@moefe/vue-aplayer';
    import {
        queryMusicList,
    } from '@/api/stat'

    export default {
        name: "musicPlayer",
        components:{
            APlayer
        },
        data(){
            return {
                audio :[]
            }
        },
        async created(){
            const res = await queryMusicList();
            this.audio = res.data;
        }
    }
</script>

<style scoped>

</style>
