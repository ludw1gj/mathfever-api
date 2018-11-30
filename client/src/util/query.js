import axios from 'axios';

/** Get categories from API. */
export const getCategories = () => {
  return axios.get('https://api-mathfever.robertjeffs.com/categories');
};

/**
 * Get category from API.
 * @param {string} categorySlug - The slug of the category to get data for.
 * @param {object | null} categories - The categories data.
 * @returns - Promise.
 */
export const getCategory = (categorySlug, categories) => {
  const getCategoryData = new Promise((resolve, reject) => {
    if (!categories) {
      axios
        .get(`https://api-mathfever.robertjeffs.com/category/${categorySlug}`)
        .then(resp => {
          resolve(resp.data);
        })
        .catch(err => {
          reject(err);
        });
    } else {
      // categories data is present, don't query again for the data.
      const category = categories.filter(category => {
        return category.slug === categorySlug;
      });
      if (category.length !== 1) {
        reject(Error('Data returned from server in redux is invalid'));
        return;
      }
      resolve(category[0]);
    }
  });
  return getCategoryData;
};

/**
 * Get calculation from API.
 * @param {string} categorySlug - The slug of the category of the calculation to get data for.
 * @param {string} calculationSlug - The slug of the calculation to get data for.
 * @param {object | null} categories - The categories data.
 * @returns - Promise.
 */
export const getCalculation = (categorySlug, calculationSlug, categories) => {
  const getCalcData = new Promise((resolve, reject) => {
    if (!categories) {
      // query for the calculation.
      axios
        .get(
          `https://api-mathfever.robertjeffs.com/calculation/${calculationSlug}`
        )
        .then(resp => {
          const calc = resp.data;
          if (calc.categorySlug !== categorySlug) {
            reject(Error('Category is invalid'));
            return;
          }
          resolve(calc);
        })
        .catch(err => {
          reject(err);
        });
    } else {
      // categories data is present, don't query again for the data.

      // filter categories.
      const category = categories.filter(category => {
        return category.slug === categorySlug;
      });
      if (category.length !== 1) {
        reject(Error('Data returned from server in redux is invalid'));
        return;
      }

      // filter calculations.
      const calculation = category[0].calculations.filter(calc => {
        return calc.slug === calculationSlug;
      });
      if (calculation.length !== 1) {
        reject(Error('Data returned from server in redux is invalid'));
        return;
      }
      resolve(calculation[0]);
    }
  });
  return getCalcData;
};

/**
 * Get calculation output from API.
 * @param {string} calculationSlug - The slug of the calculation to get data for.
 * @param {object} serializedForm - A serialize form.
 * @returns - Promise.
 */
export const processCalculation = (calculationSlug, serializedForm) => {
  return new Promise((resolve, reject) => {
    axios
      .post(
        `/api/calculation/${calculationSlug}`,
        JSON.stringify(serializedForm)
      )
      .then(resp => {
        resolve(resp.data.content);
      })
      .catch(err => {
        console.log(err.response);
        reject(err);
      });
  });
};
