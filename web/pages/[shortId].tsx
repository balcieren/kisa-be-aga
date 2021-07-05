import { GetServerSideProps, NextPage } from "next";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { api } from "../utils/api";

const ShortId = ({ shortId }) => null;

export default ShortId;

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { shortId } = context.params;

  try {
    await api.get(`/url/${shortId}/status`);
    context.res
      .writeHead(301, {
        Location: `${process.env.BASE_API_URL}/url/${shortId}`,
      })
      .end();
  } catch (error) {
    console.log("catch err:: ", error.message);
    return { notFound: true };
  }
  return {
    props: { shortId },
  };
};
