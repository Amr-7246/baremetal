#include <opencv2/opencv.hpp>
#include <iostream>
#include <vector>

// using namespace cv; //! I will not use the namespace here to avoid frustrating

int main(){
    cv::Scalar mainColor(0,255,0) ;
    //& open the webcam with validation
        cv::VideoCapture cam(0);
        if (!cam.isOpened())
        {
            std::cout << "Error: Could not open webcam." << std::endl; //! endl over /n to fkush the buffer
            return -1;
        }
    
    //& scan the screen to detict the pre-defined target color range 
        cv::Scalar lowBound(35, 40, 40);
        cv::Scalar upperBound(85, 255, 255);
        while (true)
        {
            //~ Read the cam frames + convert to HSV
            cv::Mat frame, hsvFrame, maskedFrame;
            // cam >> frame; //! the extraction operator is valid too
            flip(frame, frame, 1);
            bool frameCap = cam.read(frame);
            if (!frameCap || frame.empty()) break;
            cv::cvtColor(frame, hsvFrame, cv::COLOR_BGR2HSV); //! Here i convert the screen colors to HSV to avoid the changing shadows and lighting unreliablity

            //~ isolate the frames target color + Reduce the frames noise
            cv::inRange(hsvFrame, lowBound, upperBound, maskedFrame);
            cv::Mat kernel = cv::getStructuringElement(cv::MORPH_ELLIPSE, cv::Size(5,5));
            cv::erode(maskedFrame, maskedFrame, kernel);
            cv::dilate(maskedFrame, maskedFrame, kernel);

    //& draw a rectangle arround the target
            //~ Find the particle corners Contours (for each shap with the determined color range)
            std::vector<std::vector<cv::Point>> contours; // [ [(0,5), (..), (), ()], [(0,5), (..), (), ()], [...] ]
            //! we just care about the external boundary, so I detrmined the RETR_EXTERNAL mode
            //! the CHAIN_APPROX_SIMPLE method saves just the corners of the square (I do not need the full shap contours)
            cv::findContours(maskedFrame, contours, cv::RETR_EXTERNAL, cv::CHAIN_APPROX_SIMPLE);
            
            if (!contours.empty())
            {
            //~ Filter the bigest shap
            double maxArea = 0;
            int maxAreaIdx = -1;
            for (size_t i = 0; i < contours.size(); i++)
            {
                double thisArea = cv::contourArea(contours[i]);
                if (thisArea > maxArea){ 
                    maxArea = thisArea 
                    maxAreaIdx = i
                }
            }
            
            //~ Draw a rectangle around the shap + Draw a plus at its center + write some statistics too
            if (maxAreaIdx != -1 && maxArea > 500) //! to avoid the micro shap selection
            {
                cv::Rect box = cv::boundingRect(contours[maxAreaIdx]); //! Rect type: object with (x, y, width, height)
                cv::rectangle(frame, box, mainColor, 2); //! here we use the second version which is accept the rect over the other which is accept the corrner pointes because we do not know theme

                // int middelX = ((box.x + box.width) + box.x) / 2
                int middelX = box.x + (box.width / 2);
                int middelY = box.y + (box.height / 2);
                cv::line(frame, cv::Point(middelX-10, middelY), cv::Point(middelX+10, middelY), mainColor, 2);
                cv::line(frame, cv::Point(middelX, middelY-10), cv::Point(middelX, middelY+10), mainColor, 2);

                std::string coords = "X:" + std::to_string(middelX) + " Y:" + std::to_string(middelY);
                cv::putText(frame, "TARGET LOCKED [ACTIVE]", cv::Point(box.x, box.y - 25), cv::FONT_HERSHEY_SIMPLEX, 0.6, mainColor, 2);
                cv::putText(frame, coords, cv::Point(box.x, box.y - 10), cv::FONT_HERSHEY_PLAIN, 1, mainColor, 1);
            }
            
            } else {
                cv::putText(frame, "We do not see any thing AMR...", cv::Point(50, 50), cv::FONT_HERSHEY_SIMPLEX, 1, cv::Scalar(0,0,255), 2);
            }
            
            //~ display result and handle the ESC close key
            cv::imshow("Defense Tech Tracker - C++", frame);
            cv::imshow("Computer Vision Mask", maskedFrame);
            if(cv::waitKey(30) == 27)break; //! ESC 
        }
    cam.release();
    cv::destroyAllWindows();
    return 0;

}