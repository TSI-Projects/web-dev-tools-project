<template>
    <q-page padding>
        <div class="row q-col-gutter-md">
            <div
                v-for="ad in ads"
                class="col-6 col-xl-2 col-lg-2 col-md-3 col-sm-4"
            >
                <ad-card
                    :id="ad.id"
                    :title="ad.title"
                    :description="ad.description"
                    :price="ad.price"
                    :preview-img="ad.preview_img"
                />
            </div>
        </div>
    </q-page>
</template>

<script lang="ts" setup>
interface Ad
{
    id: string;
    title :string;
    preview_img: string;
    description: string;
    price: string;
}

const useFetchMock = (delay: number = 1000) => {
    return new Promise<Ad[]>((resolve) => {
        setTimeout(() => {
            resolve([
                {
                    id: "h182hofa8yhr-4hiofau3-ui3o231jv",
                    title: "Asus Expertbook",
                    preview_img: "https://i.ss.lv/gallery/6/1120/279831/55966116.th2.jpg",
                    description: "Pārdodu jaunu portatīvo datoru.",
                    price: "330€",
                }
            ]);
        }, delay);
    });
};

const { data: ads } = await useAsyncData<Ad[]>(() => useFetchMock());
</script>